<template>
    <div>
        <v-row justify="center">
            <h2>{{heading}}</h2>
        </v-row>
        <v-row justify="center">
            <table>
                <tr v-for="item in plants" :key="item.plant_type">
                    <td>{{item.plant_type}}</td>
                    <td>{{item.available_plants}}</td>
                </tr>
            </table>
        </v-row>
    </div>
</template>

<script>
    export default {
        name: "InfoBoxPlants",
        props: {
            heading: String,
            plants: Array
        },
        data() {
            return {

            }
        }
    }
</script>

<style scoped>
    h2{
        font-style: normal;
        font-weight: 500;
        font-size: 24px;
        line-height: 29px;
        color: white;
    }

    table{
        color: white;
    }

    td{
        padding: 0 10px;
    }
</style>