module.exports = {
  extends: ['stylelint-config-standard', 'stylelint-config-prettier'],
  // add your custom config here
  // https://stylelint.io/user-guide/configuration
  rules: {
    'no-descending-specificity': null,
    'at-rule-no-unknown': [
      true,
      {
        ignoreAtRules: [
          'function',
          'if',
          'for',
          'each',
          'include',
          'mixin',
          'content',
          'extend',
        ],
      },
    ],
    // 関数名には小文字、大文字を無視
    'function-name-case': null,
    // 汎用フォントファミリ欠落を許可
    'font-family-no-missing-generic-family-keyword': null,
    'no-duplicate-selectors': null,
    'selector-pseudo-element-no-unknown': [
      true,
      {
        ignorePseudoElements: ['v-deep'],
      },
    ],
  },
}
