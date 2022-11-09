module.exports = {
    transpileDependencies: ["vuetify"],
    css: {
        loaderOptions: {
            scss: {
                additionalData: `
                @import "@/styles/app.scss";
            `
            }
        }
    },
    devServer: {
        disableHostCheck: true
    }
};