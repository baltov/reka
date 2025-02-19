// types.ts

export interface Image {
  id: number;
  // You can add more properties here if needed, such as URL or description
  // Example:
  // url: string;
  // description?: string;
}

export interface Answer {
  id: number;
  text: string;
  correct: boolean;
  isCorrect: boolean;
}

export interface Question {
  id: number;
  text: string;
  images: Image[];
  answers: Answer[];
}