<template>
  <div>
      <v-row>
                  <h1 style="color:#789659">Enter the admin Password</h1>
              </v-row>
              <v-row justify="center">
                  <span v-for="index in 4" :key="index">
                    <div style="margin:5px;"  v-if="index <= entry.length">
                  <v-icon large color="primary">mdi-numeric-{{entry[index-1]}}-box-outline</v-icon>
                    </div>
                    <div style="margin:5px;" v-else>
                        <v-icon large color="primary">mdi-checkbox-blank-circle-outline</v-icon>
                    </div>
                  </span>
              </v-row>
              <v-row style="margin:20px 0 10px 0;">
        <v-btn style="margin:5px;" min-height="75px" min-width="125px" color="primary" x-large rounded @click="addNumber(1)" dark fab><v-icon>mdi-numeric-1-box-outline</v-icon></v-btn>
        <v-btn style="margin:5px;" min-height="75px" min-width="125px" color="primary" x-large rounded @click="addNumber(2)" dark fab><v-icon>mdi-numeric-2-box-outline</v-icon></v-btn>
        <v-btn style="margin:5px;" min-height="75px" min-width="125px" color="primary" x-large rounded @click="addNumber(3)" dark fab><v-icon>mdi-numeric-3-box-outline</v-icon></v-btn>
              </v-row>
                <v-row style="margin:20px 0 10px 0;">
        <v-btn style="margin:5px;" min-height="75px" min-width="125px" color="primary" x-large rounded @click="addNumber(4)" dark fab><v-icon>mdi-numeric-4-box-outline</v-icon></v-btn>
        <v-btn style="margin:5px;" min-height="75px" min-width="125px" color="primary" x-large rounded @click="addNumber(5)" dark fab><v-icon>mdi-numeric-5-box-outline</v-icon></v-btn>
        <v-btn style="margin:5px;" min-height="75px" min-width="125px" color="primary" x-large  rounded @click="addNumber(6)" dark fab><v-icon>mdi-numeric-6-box-outline</v-icon></v-btn>
              </v-row>
                <v-row style="margin:20px 0 10px 0;"> 
        <v-btn style="margin:5px;" min-height="75px" min-width="125px" color="primary" x-large rounded @click="addNumber(7)" dark fab><v-icon>mdi-numeric-7-box-outline</v-icon></v-btn>
        <v-btn style="margin:5px;" min-height="75px" min-width="125px" color="primary" x-large rounded @click="addNumber(8)" dark fab><v-icon>mdi-numeric-8-box-outline</v-icon></v-btn>
        <v-btn style="margin:5px;" min-height="75px" min-width="125px" color="primary" x-large rounded @click="addNumber(9)" dark fab><v-icon>mdi-numeric-9-box-outline</v-icon></v-btn>
              </v-row>
               <v-row style="margin:20px 0 10px 0;"> 
        <v-btn style="margin:5px;" min-height="75px" min-width="125px" color="primary" x-large rounded @click="clearNumber()" dark fab><v-icon>mdi-close</v-icon></v-btn>
        <v-btn style="margin:5px;" min-height="75px" min-width="125px" color="primary" x-large rounded @click="addNumber(0)" dark fab><v-icon>mdi-numeric-0-box-outline</v-icon></v-btn>
        <v-btn style="margin:5px;" min-height="75px" min-width="125px" color="primary" x-large rounded @click="removeNumber()" dark fab><v-icon>mdi-arrow-left</v-icon></v-btn>
              </v-row>
              <v-row style="margin:20px 0 10px 0;"> 
        <v-btn style="margin:5px;" min-height="75px" min-width="375px" color="primary" icon x-large rounded @click="savePW()">Save</v-btn>
              </v-row>
    
  </div>
</template>

<script>
//style="width: 290px; flex: 0 1 auto;"
import pw_json from "@/components/admin/admin.json"
export default {
  name: "PasswordSettings",
  components: {},
  data() {
    return {
        pw: pw_json.pw,
            entry: "",
            zIndex: 2,
            
    
    };
  },
  methods: {
      addNumber(entnumber){
                 this.entry = this.entry+entnumber
      
                 if (this.entry.length > 4){
                     this.entry = this.entry.slice(1)
                 }
             },
             removeNumber(){
                 this.entry = this.entry.slice(0, -1)
             },
             clearNumber(){
                 this.entry = ""
             },
             savePW(){
                const data = JSON.stringify(this.entry)
                const fs = require('fs');
                fs.writeFile('@/components/admin/admin.json', data); 

             },
   
  },
  created() {
    
  },

  computed: {
    saveActive: function() {
      return this.entry.length == 4
    }

  }
};
</script>

<style scoped>
#logo-text {
  font-family: Montserrat, sans-serif;
  font-style: normal;
  font-weight: 600;
  font-size: 36px;
  color: var(--v-primary-base);
}
</style>