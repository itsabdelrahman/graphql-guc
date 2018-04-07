import { capitalize, encrypt, decrypt } from '../utilities';

describe('capitalize()', () => {
  test('It should capitalize the first character in a string', () => {
    expect(capitalize('quiz')).toBe('Quiz');
  });
});

describe('encrypt()', () => {
  test('It should encrypt a plain string', () => {
    expect(encrypt('secret_key')('secret_data')).toBe(
      'b40834e6b6bd202b05c47d9a0adf3b08',
    );
  });
});

describe('decrypt()', () => {
  test('It should decrypt an encrypted string', () => {
    expect(decrypt('secret_key')('b40834e6b6bd202b05c47d9a0adf3b08')).toBe(
      'secret_data',
    );
  });
});
