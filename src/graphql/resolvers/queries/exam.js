import R from 'ramda';
import { getExams } from '../../../datasource';

const examResolver = async (obj, args, context) => {
  const { code } = obj;
  const { username, password } = context;

  const exams = await getExams({ username, password });
  return R.find(R.propEq('code', code))(exams);
};

export default examResolver;
