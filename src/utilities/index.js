import R from 'ramda';

export const capitalize = R.compose(
  R.join(''),
  R.over(R.lensIndex(0), R.toUpper),
);
