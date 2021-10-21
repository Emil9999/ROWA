// src/plugins/vuetify.js

import Vue from 'vue'
import Vuetify from 'vuetify/lib'
import AdminSettings from "@/components/icons/AdminSettings";

Vue.use(Vuetify)

export default new Vuetify({
    theme: {
        themes: {
            light: {
                primary: "#009966",
                secondary: "#F7F9F7",
                accent: "#E3927B",
            },
        },
        options: {
            customProperties: true,
        },
    },
    icons: {
        iconfont: 'mdiSvg',
        values: {
            adminSettings: { // name of our custom icon
                component: AdminSettings // our custom component
            },
        },
    }
})
