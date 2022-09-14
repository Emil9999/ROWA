module.exports = {
  purge: [],
  darkMode: false, // or 'media' or 'class'
  theme: {
    extend: {
      fontFamily: {
        mulish: ["MULISH"]
      },
      fontSize:{
        '4xl': ['36px', '45px'],
      },
      borderRadius: {
        'bottom-sheet': '60px'
      },
      spacing:{
        '102' : '390px',
        '135' : '33.75rem',
        '200' : '45rem',
        'farmH' : '70rem',
        'farmW' : '45rem',
      },
      dropShadow: {
        'md': '0 10px 60px rgba(100, 140, 117, 0.2)',
        
      },
      boxShadow: {
        'list' : '0px 5px 30px rgba(100, 140, 117, 0.2)'
      },
      animation: {
        errorwiggle: 'errorwiggle 1s ease-in-out infinite',
        writepulse:	'writepulse 1s cubic-bezier(0.4, 0, 0.6, 1) infinite',


      },
      keyframes: {
        errorwiggle: {
          '0%': {'background-color': 'bg-green'},
          '10%, 90%': { transform: 'translate3d(-1px, 0, 0)' },
          '20%, 80%': { transform: 'translate3d(2px, 0, 0)' },
          '30%, 50%, 70%': { transform: 'translate3d(-4px, 0, 0)' },
          '40%, 60%': { transform: 'translate3d(4px, 0, 0)' },
        },
        writepulse: {
          '0%, 50%': {'opacity': '1'},
          '25%, 75%': {'opacity': '.5'},
        }
      },

    },
    colors: {
      gradientCol: {
        DEFAULT: '#534E3B'
      },
      green: {
        DEFAULT: '#009966',
        dark: '#006e5f'
      },
      accentwhite: {
        DEFAULT: '#f1f5f2',
        light: '#f7f9f7'
      },
      almostwhite: {
        DEFAULT: '#fdfef7'
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
      red: {
        DEFAULT: '#c73e1d'
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
