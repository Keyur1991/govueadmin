module.exports = {
    root: true,

    env: {
        node: true,
    },

    extends: ["plugin:vue/essential", "eslint:recommended", "@vue/prettier"],

    parserOptions: {
        parser: 'babel-eslint',
    },

    rules: {
      'no-console': 'off',
      'no-debugger': 'off',
      'no-unused-vars': 'off',
      'vue/no-unused-vars': 'warn',
      'vue/no-side-effects-in-computed-properties': 'warn'
    },

    overrides: [{
        files: [
            '**/__tests__/*.{j,t}s?(x)',
            '**/tests/unit/**/*.spec.{j,t}s?(x)',
        ],
        env: {
            mocha: true,
        },
    }, ],

    'extends': [
      'plugin:vue/essential',
      'eslint:recommended',
      '@vue/prettier'
    ]
};