import crypto from 'crypto';

export const capitalize = str => [str[0].toUpperCase(), str.slice(1)].join('');

export const encrypt = key => data => {
  const cipher = crypto.createCipher('aes-256-cbc', key);
  return cipher.update(data, 'utf-8', 'hex').concat(cipher.final('hex'));
};

export const decrypt = key => encryptedData => {
  const decipher = crypto.createDecipher('aes-256-cbc', key);
  return decipher
    .update(encryptedData, 'hex', 'utf-8')
    .concat(decipher.final('utf-8'));
};

export const encryptWithStoredKey = encrypt(process.env.encryption_key);

export const decryptWithStoredKey = decrypt(process.env.encryption_key);
