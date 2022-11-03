<template>
  <v-form v-model="valid">
    <v-select
      :items="receivedAddressLists"
      v-validate="'required'"
      data-vv-name="received_address"
      :error-messages="errors.collect('received_address')"
      v-model="formModel.received_address"
      label="Receive address"
      auto
      required
      item-text="name"
      item-value="address"
    ></v-select>
    <v-text-field
      label="Name"
      placeholder="WAHAHA HONPO"
      name="name"
      v-validate="'required'"
      data-vv-name="name"
      :error-messages="errors.collect('name')"
      v-model="formModel.name"
      required
    ></v-text-field>
    <v-text-field
      label="Tel"
      placeholder="(###) ### - ####"
      v-validate="'required'"
      data-vv-name="tel"
      :error-messages="errors.collect('tel')"
      v-model="formModel.tel"
      mask="phone"
      required
    ></v-text-field>
    <v-text-field
      label="Deposit"
      v-validate="'required'"
      data-vv-name="deposit"
      v-model.number="formModel.deposit"
      :error-messages="errors.collect('deposit')"
      required
    ></v-text-field>
    <v-text-field
      label="Address"
      placeholder="4-2-8 Shibakoen, Minato-ku, Tokyo, 105-0011, Japan"
      v-validate="'required'"
      data-vv-name="merchant_address"
      v-model="formModel.merchant_address"
      :error-messages="errors.collect('merchant_address')"
      required
    ></v-text-field>
    <v-select
      :items="cancelableTimeList"
      v-validate="'required'"
      data-vv-name="cancelable_time"
      :error-messages="errors.collect('cancelable_time')"
      v-model="formModel.cancelable_time"
      label="Cancelable"
      auto
      required
    ></v-select>
    <div class="form-btn">
      <v-btn outlined @click="submit" color="primary">Submit</v-btn>
      <v-btn outlined @click="clear">Clear</v-btn>
    </div>
  </v-form>
</template>

<script>
import tokenContactAddresses from '@/constants/tokenContactAddresses';

export default {
  $_veeValidate: {
    validator: 'new'
  },
  data: () => ({
    formModel: {},
    valid: true,
  }),
  computed: {
    receivedAddressLists() {
      return tokenContactAddresses.TOKEN_CONTRACT_ADDRESSES
    },
    cancelableTimeList() {
      let list = []
      for(let i=1;i <= 31;i++){
        list.push(
          {
            text:`before ${i} days`,
            value: i * 60 * 60 // to minute unit
          }
        )
      }
      return list
    }
  },
  methods: {
    /**
     * received_address(string)：受け取りたいトークンのコントラクトアドレス（ERC20）
     * name(string) ：店舗名
     * tel(string)：電話番号
     * merchant_address(string)：住所
     * deposit(int)：必要デポジット金額
     * cancelable_time(timestamp)：キャンセル可能期間
     */
    submit() {
      this.$validator.validateAll();
      try {
        this.$axios.post('/api/register', {
          received_address: this.formModel.received_address,
          name: this.formModel.name,
          tel: this.formModel.tel,
          merchant_address: this.formModel.merchant_address,
          deposit: this.formModel.deposit,
          cancelable_time: this.formModel.cancelable_time,
        })
      }catch (e){
        console.log(e)
      }
    },
    clear() {
      this.formModel = {};
      this.$validator.reset();
    }
  }
};
</script>
