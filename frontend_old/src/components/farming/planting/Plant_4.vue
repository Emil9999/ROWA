<template>
<v-container>
 <v-row justify="center"> 
     <h1>Move the plants <br>to the outside</h1>
 </v-row>
  <v-row justify="center"> 
  <p>Move each plant one week up</p>
  </v-row>
  <v-row justify="center"> 
 
          
      
           <img :src="getImgUrlForMoveInstruc()">"
  
  </v-row>
  <v-row style="margin: 0 0px -60px 80px;"> 
   <p>Put the pot with the <br>seeds into the free <br>spot at week 1</p>
  </v-row>
 <v-row align-center justify="center"> 

     
        
    
       
          <img :src="getImgUrlForInsertInstruc()">
      
  </v-row>
 

    
    </v-container>
</template>

<script>

export default {
    name: "plant_4",
    components:{
   
         
    },
    props:{
      module: Number,
       
    },
    methods:{
      getImgUrlForMoveInstruc() {
            if(this.module%2 == 0){
                return require('@/assets/harvesting/MoveInstructionsRight.svg')
            } else {
                return require('@/assets/harvesting/MoveInstructionsLeft.svg') 
            }
            },
             getImgUrlForInsertInstruc() {
            if(this.module%2 == 0){
                return require('@/assets/harvesting/InsertInstructionsRight.svg')
            } else {
                return require('@/assets/harvesting/InsertInstructionsLeft.svg') 
            }
            },
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
h1{
    
      
font-family: Montserrat;
font-style: normal;
font-weight: 600;
text-align: center;

color:var(--v-primary-base)
}
</style>