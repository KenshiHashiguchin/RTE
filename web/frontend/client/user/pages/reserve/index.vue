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
                           @close="closeSettlementModal"
          ></SettlementModal>
          <CancelDialog :active="isActiveCancelModal"
                        :message="cancelMsg"
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
      cancelReservation: null,
    }
  },
  components: {CancelDialog, Breadcrumbs},
  async asyncData({app, error}) {
    let reservations = []
    try {
      const {data} = await app.$axios.get('/api/reserve_list')
      reservations = data.reservations ? data.reservations.map(v => {
        v.status = 'Loading'
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
      const statuses = [
        'None',
        'Reserved',
        'Canceled',
        'Settled',
        'Loading', // front用
      ]
      return statuses[Math.floor(Math.random() * statuses.length)]
    },
    cancelMsg() {
      return 'Can I really cancel this reservation?'
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
        console.log('item.id')
        console.log(item.id)
        const targetReservation = this.reservations.find(v => v.id = paymentId)

        const instance = this.createWeb3Instance(Web3.givenProvider)
        console.log(instance)
        const contract = await this.getContract(instance, 'trustReserve')
        console.log(contract)
        const res = await contract.methods.getReservation(paymentId).call()
        if (res) {
          targetReservation.status = res.status
        }
      } catch (error) {
        console.log(error)
      }
    },
    async execSettleReservationContract() {
      try {
        const instance = this.createWeb3Instance(Web3.givenProvider)
        const accounts = await instance.eth.getAccounts()
        const account = accounts[0]
        const contract = await this.getContract(instance)
        const res = await contract.methods.settleReservation().send({from: account})
        console.log(res)
        this.reservation = res
      } catch (error) {
        console.log(error)
      }
    },
    async execCancelContract() {
      try {
        const paymentId = this.cancelReservation.payment_id
        const instance = this.createWeb3Instance(Web3.givenProvider)
        const accounts = await instance.eth.getAccounts()
        const account = accounts[0]
        const contract = await this.getContract(instance)
        const res = await contract.methods.cancel(paymentId).send({from: account})

        console.log(res)
      } catch (error) {
        console.log(error)
      }
    },
    activeSettlementModal(item) {
      this.isActiveSettlementModal = true
      this.settlementReservation = item
    },
    closeSettlementModal() {
      this.isActiveSettlementModal = false
      this.settlementReservation = null
    },
    activeCancelModal(item) {
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
