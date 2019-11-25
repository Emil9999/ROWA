<template>
<v-container>
<v-row justify="center">
    <h1  style="color:#789659">You will harvest from Module</h1>
</v-row>
    <v-row justify="center" style="margin-top: 40px">
    <CatTree></CatTree>
     </v-row>

    
    <v-row class="info-box" justify="center"> 
      
       <v-col align-self="center" align="center"> <h3 v-text="selectedPlant"></h3> </v-col>
        <v-col> <img src="../../../assets/logo.svg" alt="" width="80px" height="auto"> </v-col>
       
   </v-row>
   

    
    </v-container>
</template>

<script>
import CatTree from "../../home/Farm/CatTree"

import axios from "axios"
export default {
    name: "harvest_2",
    components:{
         CatTree,
    },
    props:{
        selectedPlant: String
    },
    data(){
        return{
            selectedplant:this.selectedPlant,
            positionModule: null,
        }
    },
    methods:{
        getPosAndModule(){
            axios.post("http://127.0.0.1:3000/harvest/harvest",  { plantType: this.selectedplant },
          "content-type: application/json"
        ).then(result => {
                        this.positionModule = result.data
                        console.log(this.positionModule)
                    })
                    .catch(error => {
                        console.log(error)
                    })

        }
    },
    created(){
        this.getPosAndModule()
    }

    
}
</script>

<style scoped>

.info-box {
  background: #ffffff;
  border-radius: 10px;
  box-shadow: 0px 4px 4px rgba(0, 0, 0, 0.25);
  margin: 20px 90px 0 90px;
}

#button{
 font-weight: bold;
 margin: 40px;
 font-family: Montserrat;
 font-size: 24px
}
</style>