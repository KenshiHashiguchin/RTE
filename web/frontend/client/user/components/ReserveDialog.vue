<template>
  <v-dialog
      v-model="dialog"
      light
      max-width="600px"
      transition="dialog-left-transition"
  >
    <template v-slot:activator="{ on, attrs }">
      <v-btn
          x-small
          v-bind="attrs"
          v-on="on"
      >
        <slot>Reserve</slot>
      </v-btn>
    </template>
    <v-card>
      <v-card-title>
        <span class="text-h5">{{ merchant.name }}</span>
      </v-card-title>
      <v-card-text>
        <br>
        <div>
          <p><b>Required deposit amount</b>：{{ merchant.deposit }}({{ erc20Address }})
            <small>　　※You are free to choose the tokens you wish to pay.</small>
          </p>
        </div>
        <template v-if="!showPayment">
          <b>Reservation Information</b>
          <v-container>
            <v-row>
              <v-col
                  cols="6"
                  sm="6"
                  md="6"
              >
                <v-text-field
                    label="Legal first name*"
                    required
                    v-model="firstname"
                ></v-text-field>
              </v-col>
              <v-col
                  cols="6"
                  sm="6"
                  md="6"
              >
                <v-text-field
                    label="Legal last name*"
                    persistent-hint
                    required
                    v-model="surname"
                ></v-text-field>
              </v-col>
              <v-col cols="12">
                <v-text-field
                    label="Tel*"
                    required
                    type="tel"
                    v-model="phonenumber"
                ></v-text-field>
              </v-col>
              <v-col
                  cols="12"
                  sm="6"
              >
                <v-menu
                    ref="dateCalender"
                    v-model="dateCalender"
                    :return-value.sync="date"
                    transition="scale-transition"
                    offset-y
                    min-width="auto"
                >
                  <template v-slot:activator="{ on, attrs }">
                    <v-text-field
                        v-model="date"
                        label="Date of visit*"
                        prepend-icon="mdi-calendar"
                        readonly
                        v-bind="attrs"
                        v-on="on"
                    ></v-text-field>
                  </template>
                  <v-date-picker
                      v-model="date"
                      no-title
                      scrollable
                      :min="this.reservableDate"
                      @click:date="$refs.dateCalender.save(date)"
                  >
                  </v-date-picker>
                </v-menu>
              </v-col>
              <v-col
                  cols="12"
                  sm="6"
              >
                <v-menu
                    ref="timeMenu"
                    v-model="timeMenu"
                    :close-on-content-click="false"
                    :nudge-right="40"
                    :return-value.sync="time"
                    transition="scale-transition"
                    offset-y
                    max-width="290px"
                    min-width="290px"
                >
                  <template v-slot:activator="{ on, attrs }">
                    <v-text-field
                        v-model="time"
                        label="Time of visit*"
                        prepend-icon="mdi-clock-time-four-outline"
                        readonly
                        v-bind="attrs"
                        v-on="on"
                    ></v-text-field>
                  </template>
                  <v-time-picker
                      v-if="timeMenu"
                      v-model="time"
                      :allowed-minutes="allowedStep"
                      full-width
                      @click:minute="$refs.timeMenu.save(time)"
                  ></v-time-picker>
                </v-menu>
              </v-col>
              <v-col
                  cols="12"
                  sm="6"
              >
                <v-select
                    :items="[...Array(20)].map((_, i) => i+1)"
                    label="Number of visiters*"
                    required
                    v-model="number"
                ></v-select>
              </v-col>
            </v-row>
            <p v-if="isError" style="color: red">*indicates required field</p>
            <p v-else>*indicates required field</p>
          </v-container>
          <v-btn
              class="d-flex justify-center align-center"
              color="orange darken-2"
              @click="submit()"
          >
            Make a reservation
          </v-btn>
        </template>
        <template v-else>
          <Payment
              @close='showPayment = false'
              :receive-token-amount="merchant.deposit"
              :receive-token-name="merchant.name"
              :receive-address="merchant.received_address"
              :reservation="{surname: surname,firstname: firstname,phonenumber: phonenumber,number: number,date: date,time: time,}"
              :merchant="merchant"
          >Select Deposit Token
          </Payment>
        </template>
      </v-card-text>
      <v-card-actions>
        <v-row
            align="center"
            justify="space-around"
        >
        </v-row>
        <br>
        <br>
      </v-card-actions>
    </v-card>
  </v-dialog>
</template>

<script>
import tokenContactAddresses from '@/constants/tokenContactAddresses';
import Payment from "~~/client/user/components/payment/Payment";

export default {
  name: "ReserveDialog",
  components: {Payment},
  props: {
    merchant: {
      type: Object,
      default: () => {
      }
    }
  },
  data: () => ({
    dialog: false,
    dateCalender: false,
    timeMenu: false,
    showPayment: false,
    // input
    surname: '',
    firstname: '',
    phonenumber: '',
    number: '',
    date: null,
    time: null,
    isError: false,
  }),
  computed: {
    erc20Address() {
      let erc20 = Object.keys(tokenContactAddresses.TOKEN_CONTRACT_ADDRESSES).filter((key) => {
        return tokenContactAddresses.TOKEN_CONTRACT_ADDRESSES[key].address === this.merchant.received_address
      });

      if (erc20.length > 0) {
        return tokenContactAddresses.TOKEN_CONTRACT_ADDRESSES[erc20].name
      }

      return this.merchant.received_address
    },
    reservableDate() {
      let now =new Date();
      now.setDate(now.getDate() + 1);
      return now.toLocaleDateString("ja-JP", {
        year: "numeric",
        month: "2-digit",
        day: "2-digit",
      })
          .split("/")
          .join("-");
    }
  },
  methods: {
    allowedStep: m => m % 15 === 0,
    show() {
      this.showPayment = true
    },
    submit() {
      if (!this.surname || !this.firstname || !this.date || !this.time || !this.phonenumber || !this.number) {
        this.isError = true
      } else {
        this.isError = false
        this.show()
      }
    },
  }
}
</script>

<style scoped>

</style>
