import Vue from 'vue'
import Vuetify from 'vuetify/lib'
import colors from 'vuetify/lib/util/colors'


Vue.use(Vuetify)

export default new Vuetify({
  theme: {
    dark: true,
    themes: {
      // light: {
      //   primary: colors.cyan, // #00BCD4
      //   secondary: colors.cyan.lighten2, // #4DD0E1
      //   accent: colors.yellow.darken1, // #FDD835
      // },
      /**
       * @see colors.png
       */
      dark: {
        primary: '#87CFFB',
        secondary: '#4C9EED',
        accent: '#F2A45F',
        error: '#F9D6A5',
        info: '#C09A6C',
      },
    },
  },
})
