import StackDirection from './StackDirection';

export default interface GameState {
  readonly turn: number;
  readonly currentPlayerIndex: number;
  readonly log: string[];
  readonly deck: StateCard[];
  readonly deckSize: number;
  readonly score: number;
  readonly maxScore: number;
  readonly clueTokens: number;
  readonly doubleDiscard: boolean;
  readonly strikes: StateStrike[];
  readonly hands: number[][];
  readonly playStacks: number[][];
  readonly playStacksDirections: StackDirection[];
  readonly discardStacks: number[][];
  readonly clues: StateClue[];
  readonly stats: StateStats;
}

export interface StateCard {
  readonly suit: number;
  readonly rank: number;
  readonly clues: StateCardClue[];
}

export interface StateStrike {
  readonly order: number;
  readonly turn: number;
}

export interface StateClue {
  readonly type: number;
  readonly value: number;
  readonly giver: number;
  readonly target: number;
  readonly turn: number;
}

export interface StateCardClue {
  readonly type: number;
  readonly value: number;
  readonly positive: boolean;
}

export interface StateStats {
  readonly cardsGotten: number;
  readonly potentialCluesLost: number;
  readonly efficiency: number;
  readonly pace: number | null;
  readonly paceRisk: PaceRisk;
}

export type PaceRisk = 'LowRisk' | 'MediumRisk' | 'HighRisk' | 'Zero' | 'Null';
