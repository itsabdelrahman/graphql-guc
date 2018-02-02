import { student, courses, schedule, transcript } from './resolvers/queries';

const resolversMap = {
  Query: {
    student,
  },
  Student: {
    courses,
    schedule,
    transcript,
  },
};

export default resolversMap;
