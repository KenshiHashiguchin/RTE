import Vue from 'vue'
import Vuetify from 'vuetify/lib'
import colors from 'vuetify/lib/util/colors'

Vue.use(Vuetify)

export default new Vuetify({
  theme: {
    themes: {
      light: {
        primary: colors.cyan, // #00BCD4
        secondary: colors.cyan.lighten2, // #4DD0E1
        accent: colors.yellow.darken1, // #FDD835
      },
    },
  },
})
