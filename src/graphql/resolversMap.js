import { student, courses, absence, exam, schedule } from './resolvers/queries';

const resolversMap = {
  Query: {
    student,
  },
  Student: {
    courses,
    schedule,
  },
  Course: {
    absence,
    exam,
  },
};

export default resolversMap;
