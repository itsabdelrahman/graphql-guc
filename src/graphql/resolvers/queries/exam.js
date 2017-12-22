import R from 'ramda';
import { getExams } from '../../../datasource';

export default async ({ code }, args, { username, password }) =>
  R.find(R.propEq('code', code))(await getExams({ username, password }));
