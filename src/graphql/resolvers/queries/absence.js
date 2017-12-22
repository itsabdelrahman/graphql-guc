import R from 'ramda';
import { getAttendance } from '../../../datasource';

const absenceResolver = async (obj, args, context) => {
  const { code } = obj;
  const { username, password } = context;

  const attendance = await getAttendance({ username, password });
  return R.find(R.propEq('code', code))(attendance);
};

export default absenceResolver;
