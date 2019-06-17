<template>
    <div>
        <v-layout>
                <v-flex xs12 sm6 offset-sm3>
                    <v-card>
                        <v-img
                            src="https://www.jp.square-enix.com/recruit/career/contents/final_fantasy_vii_remake/img/index/final_fantasy_vii_remake-01.jpg"
                            aspect-ratio="2.75"
                        ></v-img>
                        <v-card-title>
                            <div>
                                <h3 class="headline mb-0">{{ item.Name }}</h3>
                                <div>{{ item.Description }}</div>
                                <div>{{ item.Amount }}円</div>
                                <payjp-checkout
                                    api-key="payjp-test-key-pub"
                                    text="カード情報を入力して購入"
                                    submit-text="購入確定"
                                    name-placeholder="名前"
                                    v-on:created="onTokenCreated"
                                    v-on:failed="onTokenFailed">
                                </payjp-checkout>
                            </div>
                        </v-card-title>
                    </v-card>
                    <p>{{ message }}</p>
                    <router-link to="/">HOMEへ</router-link>
                </v-flex>
            </v-layout>
    </div>
</template>

<script>
    import axios from 'axios'
    export default {
        name: 'ItemCard',
        data () {
            return {
                item: {},
                message: ''
            }
        },
        created () {
            // Take product information from the dynamic parameter specified by url
            axios.get(`http://localhost:8888/api/v1/items/${this.$route.params.id}`).then(res => {
                this.item = res.data
            })
        },
        beforeDestroy () {
            window.PayjpCheckout = null
        },
        methods: {
            // Called when the card tokenization succeeds. I will go directly to the product purchase with that Token.
            onTokenCreated: function (res) {
                console.log(res.id)
                const data = {Token: res.id}
                axios.post(`http://localhost:8888/api/v1/charge/items/${this.$route.params.id}`, data).then(res => {
                    this.message = '商品の購入が完了しました。'
                })
            },
            // Called when Tokenization fails.
            onTokenFailed: function (status, err) {
                console.log(status)
                console.log(err)
            }
        }
    }
</script>
