const polished = require('polished');
const defaultConfig = require('tailwindcss/defaultConfig');

module.exports = {
    theme: {
        extend: {
            colors: {
                'primary-100': polished.lighten(0.4, '#ff0045'),
                'primary-200': polished.lighten(0.3, '#ff0045'),
                'primary-300': polished.lighten(0.2, '#ff0045'),
                'primary-400': polished.lighten(0.1, '#ff0045'),
                'primary-500': '#ff0045',
                'primary-600': polished.darken(0.1, '#ff0045'),
                'primary-700': polished.darken(0.2, '#ff0045'),
                'primary-800': polished.darken(0.3, '#ff0045'),
                'primary-900': polished.darken(0.4, '#ff0045'),
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
