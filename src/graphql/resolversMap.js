import { student, courses, schedule } from './resolvers/queries';

const resolversMap = {
  Query: {
    student,
  },
  Student: {
    courses,
    schedule,
  },
};

export default resolversMap;
