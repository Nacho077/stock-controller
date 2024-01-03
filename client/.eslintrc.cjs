module.exports = {
  root: true,
  env: { browser: true, es2020: true },
  extends: [
    'eslint:recommended',
    'plugin:@typescript-eslint/recommended',
    'plugin:react-hooks/recommended',
  ],
  ignorePatterns: ['dist', '.eslintrc.cjs'],
  parser: '@typescript-eslint/parser',
  plugins: ['react-refresh'],
  rules: {
    'react-refresh/only-export-components': [
      'warn',
      { allowConstantExport: true },
    ],
    'no-const-assign': 'error',
    'no-fallthrough': 'error',
    'use-isnan': 'error',
    'no-irregular-whitespace': 'warn',
    'no-self-assign': 'warn',
    'no-restricted-imports': 'off',
    '@typescript-eslint/no-restricted-imports': [
      'warn',
      {
        'name': 'react-redux',
        'importNames': ['useSelector', 'useDispatch'],
        'message': 'Use typed hooks `useAppDispatch` and `useAppSelector` instead.'
      }
    ],
  },
}
