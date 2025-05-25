/** @type {import('tailwindcss').Config} */
export default {
  content: [
    './index.html',
    './src/**/*.{vue,js,ts,jsx,tsx}',
    './src/views/**/*.{vue,js,ts,jsx,tsx}',
    './src/components/**/*.{vue,js,ts,jsx,tsx}'
  ],
  theme: {
    extend: {
      zIndex: {
        '100': '100',
        '110': '110',
      }
    },
  },
  plugins: [
    require('@tailwindcss/typography'),
  ],
}

