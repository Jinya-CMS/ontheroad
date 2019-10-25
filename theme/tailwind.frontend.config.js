const polished = require('polished');
const defaultConfig = require('tailwindcss/defaultConfig');

module.exports = {
    theme: {
        extend: {
            colors: {
                'primary-100': polished.lighten(0.4, '#da9c8a'),
                'primary-200': polished.lighten(0.3, '#da9c8a'),
                'primary-300': polished.lighten(0.2, '#da9c8a'),
                'primary-400': polished.lighten(0.1, '#da9c8a'),
                'primary-500': '#da9c8a',
                'primary-600': polished.darken(0.1, '#da9c8a'),
                'primary-700': polished.darken(0.2, '#da9c8a'),
                'primary-800': polished.darken(0.3, '#da9c8a'),
                'primary-900': polished.darken(0.4, '#da9c8a'),
                'secondary-100': polished.lighten(0.4, '#8ac8da'),
                'secondary-200': polished.lighten(0.3, '#8ac8da'),
                'secondary-300': polished.lighten(0.2, '#8ac8da'),
                'secondary-400': polished.lighten(0.1, '#8ac8da'),
                'secondary-500': '#8ac8da',
                'secondary-600': polished.darken(0.1, '#8ac8da'),
                'secondary-700': polished.darken(0.2, '#8ac8da'),
                'secondary-800': polished.darken(0.3, '#8ac8da'),
                'secondary-900': polished.darken(0.4, '#8ac8da'),
                'gray-100': polished.lighten(0.4, '#adb5bd'),
                'gray-200': polished.lighten(0.3, '#adb5bd'),
                'gray-300': polished.lighten(0.2, '#adb5bd'),
                'gray-400': polished.lighten(0.1, '#adb5bd'),
                'gray-500': '#adb5bd',
                'gray-600': polished.darken(0.1, '#adb5bd'),
                'gray-700': polished.darken(0.2, '#adb5bd'),
                'gray-800': polished.darken(0.3, '#adb5bd'),
                'gray-900': polished.darken(0.4, '#adb5bd'),
            }
        }
    },
    plugins: [],
};
