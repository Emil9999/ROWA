module.exports = {
  purge: [],
  darkMode: false, // or 'media' or 'class'
  theme: {
    extend: {
      spacing:{
        '135' : '33.75rem',
        'farmH' : '70rem',
        'farmW' : '45rem',
      },
      dropShadow: {
        'md': '0px 5px 30px rgba(100, 140, 117, 0.2)'
      },
      fontFamily: {
        'sans': ['Mulish', 'sans-serif']
    },
    },
    colors: {
      green: {
        DEFAULT: '#009966',
        dark: '#006e5f'
      },
      accentwhite: {
        DEFAULT: '#f1f5f2',
        light: '#f7f9f7'
      },
      grey: {
        DEFAULT: '#c4c4c4',
        dark: '#828282'
      },
      brown: {
        DEFAULT: '#a0918d',
      },
      white: {
        DEFAULT: '#ffffff',
        
      },

      brownred: {
        DEFAULT: '#935A4A'
      }
    },
    
  },
  variants: {
    extend: {
      bg: ['disabled'],
      text: ['disabled'],
    },
  },
  plugins: [],
}
