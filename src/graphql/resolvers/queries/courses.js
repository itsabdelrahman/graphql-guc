import R from 'ramda';
import { requestCoursework } from '../../../datasource/network';
import {
  parseCourses,
  parseCoursework,
  parseMidterms,
} from '../../../datasource/parser';

const augmentCoursework = coursework => course =>
  R.assoc('coursework', R.filter(R.propEq('code', course.code))(coursework))(
    course,
  );

const augmentMidterm = midterms => course =>
  R.assoc('midterm', R.find(R.propEq('code', course.code))(midterms))(course);

const coursesResolver = async (obj, args, context) => {
  const { username, password } = context;
  const response = await requestCoursework({ username, password });

  const courses = parseCourses(response);
  const coursework = parseCoursework(response);
  const midterms = parseMidterms(response);

  return R.map(R.pipe(augmentCoursework(coursework), augmentMidterm(midterms)))(
    courses,
  );
};

export default coursesResolver;
