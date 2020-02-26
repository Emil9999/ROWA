<template>
    <div id="Admin">
<v-btn dark fab color="white" :to="{name:'Home'}">
                <v-icon color="primary">mdi-close</v-icon>
            </v-btn>

        <h1>Plan your event:</h1>
      <v-row justify="space-around" align="center">
        <v-col style="width: 290px; flex: 0 1 auto;">
          <h2>Light On:</h2>
          <h3>Currently: {{current_time.time_on}}</h3>
          <v-time-picker v-model="start" :max="end"></v-time-picker>
        </v-col>
        <v-col style="width: 290px; flex: 0 1 auto;">
          <h2>Light Off: </h2>
          <h3>Currently: {{current_time.time_off}}</h3>
          <v-time-picker v-model="end" :min="start"></v-time-picker>
        </v-col>
      </v-row>

     <v-btn id="button"   :disabled="TimeDisable" @click="sendLightTimes()" rounded color="accent" height="75" width="360">
           Save
            <v-icon right dark>mdi-checkbox-marked-circle</v-icon>
          </v-btn>
     <v-switch v-model="current_time.state" @change="changeLight()">
        <template v-slot:label>
        Light Switch
        </template>
      </v-switch>
    
    
    </div>
</template>


<script>
    import axios from "axios"

    export default {
        name: "AdminSettings",
        components: {
           
        },
         data(){
             return{
            lightStatus: false,
            start: null,
            end: null,
            current_time: {
                    time_on: null,
                    time_off: null,
                    state: null
                },
             };
         },
         methods: {
              getLightTimes: function () {
                axios.get("http://127.0.0.1:3000/adminSettings/get-light")
                    .then(result => {
                        this.current_time.time_on = result.data.time_on
                        this.current_time.time_off = result.data.time_off
                        this.current_time.state = !!result.data.state
                        /*if (result.data.state == 1){
                            this.current_time.state =  true
                        }else{
                            this.current_time.state  = false
                        }
                        */
                    })
                    .catch(error => {
                        console.log(error)
                    })
            },

            sendLightTimes:function () {
                axios.post("http://127.0.0.1:3000/adminSettings/insert-light",
                    {time_on: this.start, time_off: this.end},
                    "content-type: application/json")
                    .then()
                    .catch(error => {
                        console.log(error);
                    });

            },
             changeLight:function () {
                console.log(this.time_on)
                axios.post("http://127.0.0.1:3000/adminSettings/changelight",
                    {state: this.current_time.state| 0},
                    "content-type: application/json")
                    .then()
                    .catch(error => {
                        console.log(error);
                    }); 

            },

         },
         created() {
             this.getLightTimes()
         },
         
         computed: {
         TimeDisable: function() {
            if((this.start == null) || (this.end == null)){
            return true
            }
            else{
            return false
            }
        }
        


         }
    };
</script>

<style scoped>
    h1,
    h2 {
        font-weight: normal;
    }

    ul {
        list-style-type: none;
        padding: 0;
    }

    li {
        display: inline-block;
        margin: 0 10px;
    }

    a {
        color: #42b983;
    }

    span {
        color: var(--v-primary-base);
        font-style: normal;
        font-weight: normal;
        font-size: 14px;
    }

    .no-hover:hover {
        background-color: transparent;
        text-decoration: none;
    }
</style>