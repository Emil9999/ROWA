// src/plugins/vuetify.js

import Vue from 'vue'
import Vuetify from 'vuetify/lib'
import AdminSettings from "@/components/icons/AdminSettings";

Vue.use(Vuetify)

export default new Vuetify({
    theme: {
        themes: {
            light: {
                primary: "#789659",
                secondary: "#F4F0E6",
                accent: "#E3927B",
            },
        },
        options: {
            customProperties: true,
        },
    },
    icons: {
        values: {
            adminSettings: { // name of our custom icon
                component: AdminSettings // our custom component
            },
        },
    }
})
