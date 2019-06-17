<template>
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
                    </div>
                </v-card-title>
            </v-card>
        </v-flex>
    </v-layout>
</template>

<script>
    export default {
        name: 'ItemCard',
        props: [
            'item'
        ]
    }
</script>