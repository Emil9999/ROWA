<template>
    <div id="Admin">
        <AdminTopRow v-bind:headtext="'Change Plant Types'" v-bind:prevPage="'AdminMenu'"></AdminTopRow>
       
       
     <v-row justify="center"  margin="auto" padding="auto" v-for="item in plantTypes" :key="item.plant_type">
             <table>
               <tr>
                   <td>Module {{item.typemodule}}:</td>
                    <td>{{item.typename}}</td>
                 </tr>
                 <tr>
                       <td><img :src="getImgUrl(item.typename)" alt="" width="60px" height="auto"></td>
               
                     </tr>
             </table>
        <v-col cols="4">
                <v-select
                rounded
            v-model="item.typename"
            :items="Types"
            :menu-props="{ maxHeight: '400' }"
            label="Select"
          ></v-select>
        </v-col>
          <v-btn @click="sendModuleChanges(item.typemodule, item.typename)">Save
               <v-icon right dark>mdi-checkbox-marked-circle</v-icon>
          </v-btn>
     </v-row>
    </div>
</template>


<script>
    
    import AdminTopRow from "@/components/admin/AdminTopRow.vue"
   import axios from "axios"
    export default {
        name: "ModuleType",
        components: {
            AdminTopRow,
            
           
        },
         data(){
             return{

                 plantTypes: null,
                 Types: ["Basil"],

             };
         },
         methods: {
              getImgUrl(pic) {
                return require('@/assets/harvesting/plants/'+pic+".png")
            },
             sendModuleChanges: function(Numbe, Plant){
                 axios.post("http://127.0.0.1:3000/adminSettings/insertmodule-change",
                 {typename: Plant, typemodule: Numbe},
                 "content-type: application/json")
                 .then()
                 .catch(error => {
                     console.log(error)
                 })

             }, 
             getPlantTypes: function() {
                 axios.get("http://127.0.0.1:3000/adminSettings/get-types")
                 .then(result => {
                            this.plantTypes = result.data
                            result.data.forEach(element => { if (!this.Types.includes(element.typename)){this.Types.push(element.typename)}
                                
                            });

                            
                 })
                 .catch(error => {
                     console.log(error)
                 })
             }
              


         },
         created(){
          this.getPlantTypes()
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