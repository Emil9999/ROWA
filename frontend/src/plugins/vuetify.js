import Vue from 'vue';
import Vuetify from 'vuetify/lib/framework';

Vue.use(Vuetify);

export default new Vuetify({
    theme: {
        themes: {
          light: {
            primary: "#009966", // Susteyn green
            secondary: "#009966", // #FFCDD2
            accent: "#009966", // #3F51B5
            warning: "#E3927B",
            paragraph: "#A0918D",
          },
        },
      },
    
});
