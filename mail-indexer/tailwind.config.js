/** @type {import('tailwindcss').Config} */
module.exports = {
  purge: ['./index.html','./src/**/*.{vue,js,ts,jsx,tsx}'],
  content: [],
  theme: {
    fontFamily: {
      sans: [
        '"Inter var", sans-serif',
        {
          fontFeatureSettings: '"cv11", "ss01"',
          fontVariationSettings: '"opsz" 32'
        },
      ],
    },
    extend: {
      height: {
        '256': '48rem',
      },
      colors:{
        lightblue:"#8ECAE6",
        blue: "#219EBC",
        darkblue:"#023047",
        yellow: "#FFB703",
        orange: "#FB8500"
      }
    },
  },
  plugins: [],
}

