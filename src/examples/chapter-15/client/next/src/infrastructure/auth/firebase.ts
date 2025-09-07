import { initializeApp, FirebaseApp, FirebaseOptions } from 'firebase/app';
import { 
  getAuth, 
  Auth, 
  signInAnonymously, 
  onAuthStateChanged,
  User 
} from 'firebase/auth';
import { AuthError } from '../../domain/models';

// Check if we're in development mode
const isDevelopment = process.env.NODE_ENV === 'development';

// Firebase configuration from environment variables
const firebaseConfig: FirebaseOptions = {
  apiKey: process.env.NEXT_PUBLIC_FIREBASE_API_KEY!,
  authDomain: process.env.NEXT_PUBLIC_FIREBASE_AUTH_DOMAIN!,
  projectId: process.env.NEXT_PUBLIC_FIREBASE_PROJECT_ID!,
  storageBucket: process.env.NEXT_PUBLIC_FIREBASE_STORAGE_BUCKET,
  messagingSenderId: process.env.NEXT_PUBLIC_FIREBASE_MESSAGING_SENDER_ID,
  appId: process.env.NEXT_PUBLIC_FIREBASE_APP_ID!,
};

// Singleton instances
let app: FirebaseApp | null = null;
let auth: Auth | null = null;
let currentUser: User | null = null;
let authInitPromise: Promise<User> | null = null;

/**
 * Initialize Firebase app and auth
 */
function initializeFirebase(): { app: FirebaseApp; auth: Auth } {
  if (!app) {
    // Only initialize Firebase if we're not in development mode
    if (!isDevelopment) {
      app = initializeApp(firebaseConfig);
      auth = getAuth(app);
    }
  }
  return { app: app!, auth: auth! };
}

/**
 * Initialize Firebase auth and sign in anonymously
 * Returns a promise that resolves when auth is ready
 */
function initializeAuth(): Promise<User> {
  if (authInitPromise) {
    return authInitPromise;
  }

  // In development mode, skip Firebase
  if (isDevelopment) {
    console.warn('Using development mode without Firebase authentication');
    // Create a mock user object for development
    const mockUser = {
      uid: 'dev-user-id',
      getIdToken: async () => 'dev-token-12345',
    } as unknown as User;
    currentUser = mockUser;
    authInitPromise = Promise.resolve(mockUser);
    return authInitPromise;
  }

  const { auth } = initializeFirebase();

  authInitPromise = new Promise((resolve, reject) => {
    // Listen for auth state changes
    const unsubscribe = onAuthStateChanged(auth, async (user) => {
      if (user) {
        currentUser = user;
        unsubscribe();
        resolve(user);
      } else {
        try {
          // Sign in anonymously
          const credential = await signInAnonymously(auth);
          currentUser = credential.user;
          unsubscribe();
          resolve(credential.user);
        } catch (error) {
          unsubscribe();
          reject(new AuthError(
            `Failed to sign in anonymously: ${error instanceof Error ? error.message : 'Unknown error'}`
          ));
        }
      }
    });
  });

  return authInitPromise;
}

/**
 * Get the current Firebase ID token
 * Automatically signs in anonymously if not already authenticated
 * 
 * @returns Promise resolving to the ID token
 * @throws {AuthError} If authentication fails
 */
export async function getIdToken(): Promise<string> {
  try {
    // Ensure we're authenticated
    const user = await initializeAuth();
    
    // Get fresh ID token
    const token = await user.getIdToken();
    return token;
  } catch (error) {
    if (error instanceof AuthError) {
      throw error;
    }
    throw new AuthError(
      `Failed to get ID token: ${error instanceof Error ? error.message : 'Unknown error'}`
    );
  }
}

/**
 * Get the current user (if authenticated)
 * 
 * @returns The current user or null
 */
export function getCurrentUser(): User | null {
  return currentUser;
}

/**
 * Sign out the current user
 * Note: This will trigger a new anonymous sign-in on the next request
 */
export async function signOut(): Promise<void> {
  // In development mode, just reset
  if (isDevelopment) {
    currentUser = null;
    authInitPromise = null;
    return;
  }

  const { auth } = initializeFirebase();
  await auth.signOut();
  currentUser = null;
  authInitPromise = null;
}