// src/plugins/vuetify.js

import Vue from 'vue'
import Vuetify from 'vuetify/lib'

Vue.use(Vuetify)

export default new Vuetify({
    theme: {
        themes: {
            light: {
                primary: "#789659",
                secondary: colors.red.lighten4,
                accent: "#E3927B",
            },
        },
    },
})
