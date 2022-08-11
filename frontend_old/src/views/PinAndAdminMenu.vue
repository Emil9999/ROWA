<template>
    <div id="Admin">
        <AdminTopRow style="z-index: 9999" v-bind:headtext="'AdminPanel'" v-bind:prevPage="'Home'"></AdminTopRow>
       <div>
        <v-row justify="center" margin="auto" padding="auto" >
           <v-btn class="info-box" height="75px" width="500px" :to="{name:'AdminSettings'}">
           <h3>Admin Settings</h3> 
            </v-btn>
         </v-row>
      
        <v-row justify="center" margin="auto" padding="auto" >
           <v-btn class="info-box" height="75px" width="500px" :to="{name:'TypeChanger'}">
           <h3>Change Module Plant</h3> 
            </v-btn>
         </v-row>

         <v-row justify="center" margin="auto" padding="auto" >
           <v-btn class="info-box" height="75px" width="500px" :to="{name:'BatchFarming'}">
           <h3>Batch Farming</h3> 
            </v-btn>
         </v-row>

         <v-row justify="center" margin="auto" padding="auto" >
           <v-btn class="info-box" height="75px" width="500px" :to="{name:'RealityCheck'}">
           <h3>Reality Check</h3> 
            </v-btn>
         </v-row>

         <v-row justify="center" margin="auto" padding="auto" >
           <v-btn class="info-box" height="75px" width="500px" @click="lockAdmin()" icon>
           <h3>Lock Admin</h3> 
           <v-icon>mdi-lock</v-icon>
            </v-btn>
         </v-row>
      
          <v-overlay  :z-index="zIndex"  :absolute="true" :value="admin_active" opacity="1" color="secondary">
             <v-row>
                  <h1 style="color:#009966">Enter the admin Password</h1>
              </v-row>
              <v-row justify="center">
                  <span v-for="index in 4" :key="index">
                    <div style="margin:5px;"  v-if="index <= entry.length">
                  <v-icon large color="primary">mdi-checkbox-blank-circle</v-icon>
                    </div>
                    <div style="margin:5px;" v-else>
                        <v-icon large color="primary">mdi-checkbox-blank-circle-outline</v-icon>
                    </div>
                  </span>
              </v-row>
              <v-row style="margin:20px 0 10px 0;">
        <v-btn style="margin:5px;" min-height="75px" min-width="125px" color="primary" x-large rounded @click="addNumber(1)" dark fab><v-icon size="3.5rem">mdi-numeric-1</v-icon></v-btn>
        <v-btn style="margin:5px;" min-height="75px" min-width="125px" color="primary" x-large rounded @click="addNumber(2)" dark fab><v-icon size="3.5rem">mdi-numeric-2</v-icon></v-btn>
        <v-btn style="margin:5px;" min-height="75px" min-width="125px" color="primary" x-large rounded @click="addNumber(3)" dark fab><v-icon size="3.5rem">mdi-numeric-3</v-icon></v-btn>
              </v-row>
                <v-row style="margin:20px 0 10px 0;">
        <v-btn style="margin:5px;" min-height="75px" min-width="125px" color="primary" x-large rounded @click="addNumber(4)" dark fab><v-icon size="3.5rem">mdi-numeric-4</v-icon></v-btn>
        <v-btn style="margin:5px;" min-height="75px" min-width="125px" color="primary" x-large rounded @click="addNumber(5)" dark fab><v-icon size="3.5rem">mdi-numeric-5</v-icon></v-btn>
        <v-btn style="margin:5px;" min-height="75px" min-width="125px" color="primary" x-large  rounded @click="addNumber(6)" dark fab><v-icon size="3.5rem">mdi-numeric-6</v-icon></v-btn>
              </v-row>
                <v-row style="margin:20px 0 10px 0;"> 
        <v-btn style="margin:5px;" min-height="75px" min-width="125px" color="primary" x-large rounded @click="addNumber(7)" dark fab><v-icon size="3.5rem">mdi-numeric-7</v-icon></v-btn>
        <v-btn style="margin:5px;" min-height="75px" min-width="125px" color="primary" x-large rounded @click="addNumber(8)" dark fab><v-icon size="3.5rem">mdi-numeric-8</v-icon></v-btn>
        <v-btn style="margin:5px;" min-height="75px" min-width="125px" color="primary" x-large rounded @click="addNumber(9)" dark fab><v-icon size="3.5rem">mdi-numeric-9</v-icon></v-btn>
              </v-row>
               <v-row style="margin:20px 0 10px 0;"> 
        <v-btn style="margin:5px;" min-height="75px" min-width="125px" color="primary" x-large rounded @click="clearNumber()" dark fab><v-icon size="3.5rem">mdi-close</v-icon></v-btn>
        <v-btn style="margin:5px;" min-height="75px" min-width="125px" color="primary" x-large rounded @click="addNumber(0)" dark fab><v-icon size="3.5rem">mdi-numeric-0</v-icon></v-btn>
        <v-btn style="margin:5px;" min-height="75px" min-width="125px" color="primary" x-large rounded @click="removeNumber()" dark fab><v-icon size="3.5rem">mdi-arrow-left</v-icon></v-btn>
              </v-row>
              <v-row style="margin:20px 0 10px 0;"> 
        <v-btn style="margin:5px;" min-height="75px" min-width="375px" color="primary" icon x-large rounded :to="{name:'Home'}">Home</v-btn>
              </v-row>
    </v-overlay>

      </div>   
     
    </div>
</template>


<script>
    
    import AdminTopRow from "@/components/admin/AdminTopRow.vue"
     import {mapState, mapActions} from "vuex"
   import pw_json from "@/components/admin/admin.json"
    export default {
        name: "AdminMenu",
        components: {
            AdminTopRow,
            
           
        },
         data(){
             return{
            pw: pw_json.pw,
            entry: "",
            zIndex: 2,
             };
         },
         methods: {
             addNumber(entnumber){
                 this.entry = this.entry+entnumber
                 if (this.entry == this.pw){
                     this.change_admin_state()
                 }
                 if (this.entry.length == 4){
                     this.entry = ""
                 }
             },
             removeNumber(){
                 this.entry = this.entry.slice(0, -1)
             },
             clearNumber(){
                 this.entry = ""
             },
             lockAdmin(){
                this.change_admin_state()
                this.$router.push({name: 'Home'})
                 
             },
              
            ...mapActions(["change_admin_state"])

         },
          computed: {
            ...mapState(["admin_active"])
        },
        
        
        


         
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

    .info-box {
  background: #ffffff;
  border-radius: 10px;
  box-shadow: 0px 4px 4px rgba(0, 0, 0, 0.25);
  margin: 40px 50px 0 50px;
}
</style>