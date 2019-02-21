module.exports = {
  root: true,
  env: {
    browser: true,
    node: true
  },
  parserOptions: {
    parser: 'babel-eslint'
  },
  extends: [
    '@nuxtjs',
    'plugin:prettier/recommended'
  ],
  plugins: [
    'prettier'
  ],
  // add your custom rules here
  rules: {
    'no-unused-vars':'off',
    'camelcase':'off',
    'no-console': 'off',
    'vue/no-unused-components':'off',
    'no-undef':'off'
  }
}
