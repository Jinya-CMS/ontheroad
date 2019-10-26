const gulp = require('gulp');
const postcss = require('gulp-postcss');

function frontend() {
    return gulp.src('src/frontend.css')
        .pipe(postcss([
            require('tailwindcss')('tailwind.frontend.config.js'),
            require('autoprefixer'),
        ]))
        .pipe(gulp.dest('build/'));
}

function admin() {
    return gulp.src('src/admin.css')
        .pipe(postcss([
            require('tailwindcss'),
            require('autoprefixer'),
        ]))
        .pipe(gulp.dest('build/'));
}

exports.frontend = frontend;
exports.admin = admin;
exports.default = gulp.parallel(frontend, admin);