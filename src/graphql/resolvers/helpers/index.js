import R from 'ramda';

export const filterBy = (key, value) => array => {
  if (!value) return array;
  return R.filter(R.propEq(key, value))(array);
};
