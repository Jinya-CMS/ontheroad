const defaultConfig = require('tailwindcss/defaultConfig');

module.exports = {
    theme: {
        extend: {
            colors: {
                'primary-100': '#edcec6',
                'primary-200': '#e8c2b7',
                'primary-300': '#e3b5a8',
                'primary-400': '#dfa999',
                'primary-500': '#da9c8a',
                'primary-600': '#d58f7b',
                'primary-700': '#d1836c',
                'primary-800': '#cc765d',
                'primary-900': '#c76a4e',
                'secondary-100': '#c6e4ed',
                'secondary-200': '#b7dde8',
                'secondary-300': '#a8d6e3',
                'secondary-400': '#99cfdf',
                'secondary-500': '#8ac8da',
                'secondary-600': '#7bc1d5',
                'secondary-700': '#6cbad1',
                'secondary-800': '#5db3cc',
                'secondary-900': '#4eacc7',
                'gray-100': '#f8f9fa',
                'gray-200': '#e9ecef',
                'gray-300': '#dee2e6',
                'gray-400': '#ced4da',
                'gray-500': '#adb5bd',
                'gray-600': '#868e96',
                'gray-700': '#495057',
                'gray-800': '#343a40',
                'gray-900': '#212529',
            }
        }
    },
    variants: {
        backgroundColor: [...defaultConfig.variants.backgroundColor, 'odd'],
    },
    plugins: [],
};
