module.exports = {
  root: true,
  env: {
    browser: true,
    node: true,
    'jest/globals': true
  },
  parserOptions: {
    parser: 'babel-eslint'
  },
  extends: [
    'plugin:vue/recommended',
    'prettier'

  ],
  plugins: [
    'prettier',
    'jest'
  ],
  settings: {
    'import/resolver': {
      node: { extensions: ['.js', '.mjs'] }
    }
  },
  // add your custom rules here
  rules: {
    'no-unused-vars':'off',
    'camelcase':'off',
    'no-console': 'off',
    'vue/no-unused-components':'off',
    'no-undef':'off',
    'vue/no-v-html':'off',
    'vue/require-prop-types':'off'
  }
}
