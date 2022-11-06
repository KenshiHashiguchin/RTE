<template>
  <div id="pageReserveIndex">
    <page-header></page-header>
    <v-main>
      <v-container fluid>
        <breadcrumbs/>
        <h1>Reservations</h1>
        <v-layout>
          <v-flex text-xs-center>
            <ReserveTable :reserves="reservations"
                          @settlement="activeSettlementModal"
                          @cancel="activeCancelModal"
            >
            </ReserveTable>
          </v-flex>
          <SettlementModal
            :reservation="settlementReservation"
            :active="isActiveSettlementModal"
            :step="settlementContractStep"
            @back="backSettleReservationStep"
            @next="execSettleReservationContract"
                           @close="closeSettlementModal"
          ></SettlementModal>
          <CancelDialog :active="isActiveCancelModal"
                        :message="cancelMsg"
                        :step="cancelContractStep"
                        @next="execCancelContract"
                        @close="closeCancelModal"
          ></CancelDialog>
        </v-layout>
      </v-container>
    </v-main>
  </div>
</template>

<script>
import Breadcrumbs from "~/components/common/Breadcrumbs";
import Web3 from "web3";
import web3Mixin from "~/mixins/web3Mixin";
import CancelDialog from "~/components/common/CancelDialog";

export default {
  name: "ReserveIndex",
  mixins: [web3Mixin],
  data() {
    return {
      reservations: [],
      isActiveSettlementModal: false,
      isActiveCancelModal: false,
      settlementReservation: null,
      settlementContractStep: 1,
      cancelReservation: null,
      cancelContractStep: 1,
    }
  },
  components: {CancelDialog, Breadcrumbs},
  async asyncData({app, error}) {
    let reservations = []
    try {
      const {data} = await app.$axios.get('/api/reserve_list')
      reservations = data.reservations ? data.reservations.map(v => {
        v.status = 'Loading...'
        if (v.merchant) {
          v.merchant.cancelable_days = `before ${Math.floor(v.merchant.cancelable_time / 60 / 60)} days`
        }
        return v
      }) : []
    } catch (e) {
      error({
        statusCode: e.response.status,
        message: e.response.data.message,
      })
    }
    return {
      reservations
    }
  },
  computed: {
    statues() {
      return [
        {label: 'None', value: "0"},
        {label: 'Reserved',value: "1"},
        {label: 'Canceled',value: "2"},
        {label: 'Settled', value: "3"},
        {label: 'Loading...', value: null} // front用
      ]
    },
    cancelMsg() {
      if(this.cancelContractStep === 1){
        return 'Can I really cancel this reservation?'
      }else if(this.cancelContractStep === 2) {
        return 'Processing...'
      }else if(this.cancelContractStep === 3) {
        return ''
      }else if(this.cancelContractStep === 4) {
        return 'Cancel Committed'
      }else if(this.cancelContractStep === 9) {
        return 'Cancel Failed, Please Retry After Few Minute'
      }
    }
  },
  mounted() {
    this.reservations.forEach(v => {
      this.getReservationContract(v)
    })
  },
  methods: {
    async getReservationContract(item) {
      try {
        const paymentId = item.payment_id
        const targetReservation = this.reservations.find(v => v.payment_id === paymentId)
        const instance = this.createWeb3Instance(Web3.givenProvider)
        const contract = await this.getContract(instance, 'trustReserve')
        const res = await contract.methods.getReservation(paymentId).call()
        if (res) {
          console.log(res)
          const status = this.statues.find(v => v.value === res.status)
          targetReservation.status = status.label
          targetReservation.depositAmount = res.depositAmount
        }
      } catch (error) {
        console.log(error)
      }
    },
    async execSettleReservationContract(nextStep) {
      console.log('---nextStep---')
      console.log(nextStep)
      if(nextStep === 'inputSendAmountStep') {
        this.settlementContractStep = 2
      }else if(nextStep === 'commit'){
        this.settlementContractStep = 3
      }else if(nextStep === 'complete'){
        this.settlementContractStep = 4
        this.reservations.find(v => v.payment_id === this.settlementReservation.payment_id).status = 'Settled'
      }
      // try {
      //   const instance = this.createWeb3Instance(Web3.givenProvider)
      //   const accounts = await instance.eth.getAccounts()
      //   const account = accounts[0]
      //   const contract = await this.getContract(instance)
      //   const res = await contract.methods.settleReservation().send({from: account})
      //   console.log(res)
      //   this.reservation = res
      // } catch (error) {
      //   console.log(error)
      // }
    },
    backSettleReservationStep() {
      this.settlementContractStep = this.settlementContractStep - 1
    },
    async execCancelContract() {
      try {
        this.cancelContractStep = 2
        const paymentId = this.cancelReservation.payment_id
        const instance = this.createWeb3Instance(Web3.givenProvider)
        const accounts = await instance.eth.getAccounts()
        const account = accounts[0]
        const contract = await this.getContract(instance)
        await contract.methods.cancel(paymentId).send({from: account})
          .then((result) => {
            console.log(result)
            this.cancelContractStep = 3
            this.reservations.find(v => v.payment_id === this.cancelReservation.payment_id).status = 'Canceled'
        }).catch((err) => {
          console.log("Error!", err)
          this.cancelContractStep = 9
        })
        this.cancelReservation = null

      } catch (error) {
        console.log(error)
      }
    },
    activeSettlementModal(item) {
      this.settlementContractStep = 1
      this.isActiveSettlementModal = true
      this.settlementReservation = item
    },
    closeSettlementModal() {
      this.isActiveSettlementModal = false
      this.settlementReservation = null
    },
    activeCancelModal(item) {
      this.cancelContractStep = 1
      this.isActiveCancelModal = true
      this.cancelReservation = item
    },
    closeCancelModal(item) {
      this.isActiveCancelModal = false
      this.cancelReservation = null
    }
  }
}
</script>
