import { getSchedule } from '../../../datasource';

const schedule = async (obj, args, context) => {
  const { username, password } = context;
  return getSchedule({ username, password });
};

export default schedule;
