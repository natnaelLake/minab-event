// utils/authUtils.js

import bcrypt from 'bcryptjs';

/**
 * Hash a plain-text password
 * @param {string} password - The plain-text password
 * @returns {Promise<string>} - The hashed password
 */
export async function hashPassword(password) {
  try {
    const saltRounds = 10; // You can adjust the cost factor (default is 10)
    const hashedPassword = await bcrypt.hash(password, saltRounds);
    return hashedPassword;
  } catch (error) {
    console.error("Error hashing password:", error);
    throw new Error("Failed to hash password");
  }
}

/**
 * Compare a plain-text password with a hashed password
 * @param {string} plainPassword - The plain-text password
 * @param {string} hashedPassword - The hashed password stored in the database
 * @returns {Promise<boolean>} - Whether the passwords match or not
 */
export async function comparePasswords(plainPassword, hashedPassword) {
  try {
    const isMatch = await bcrypt.compare(plainPassword, hashedPassword);
    return isMatch;
  } catch (error) {
    console.error("Error comparing passwords:", error);
    return false; // Return false if an error occurs during comparison
  }
}
