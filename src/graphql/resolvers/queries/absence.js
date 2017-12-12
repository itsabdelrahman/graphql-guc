import R from 'ramda';
import { getAttendance } from '../../../datasource';

const absenceResolver = async (obj, args, context) => {
  const { code } = obj;
  const { username, password } = context;

  // return R.pipe(R.find(R.propEq('code', code)), R.pick(['level', 'severity']))(
  //   await getAttendance({ username, password }),
  // );

  return {};
};

export default absenceResolver;
