import {
  RecipeRequestDomain,
  RecipeResponseDomain,
  ImageRequestDomain,
  ImageResponseDomain,
  EvaluateRequestDomain,
  EvaluateResponseDomain,
} from '../../../domain/models';
import {
  RecipeRequestDTO,
  RecipeResponseDTO,
} from '../dto/recipe';
import {
  ImageRequestDTO,
  ImageResponseDTO,
} from '../dto/image';
import {
  EvaluateRequestDTO,
  EvaluateResponseDTO,
} from '../dto/evaluate';

// Recipe flow mappers
export function mapRecipeRequestToDTO(domain: RecipeRequestDomain): RecipeRequestDTO {
  return {
    ingredients: domain.ingredients,
  };
}

export function mapRecipeResponseToDomain(dto: RecipeResponseDTO): RecipeResponseDomain {
  return {
    type: dto.type as 'content' | 'done' | 'error',
    content: dto.content,
    error: dto.error,
  };
}

// Image flow mappers
export function mapImageRequestToDTO(domain: ImageRequestDomain): ImageRequestDTO {
  return {
    dishName: domain.dishName,
    description: domain.description,
  };
}

export function mapImageResponseToDomain(dto: ImageResponseDTO): ImageResponseDomain {
  return {
    success: dto.success,
    imageUrl: dto.imageUrl,
    dishName: dto.dishName,
    error: dto.error,
  };
}

// Evaluate flow mappers
export function mapEvaluateRequestToDTO(domain: EvaluateRequestDomain): EvaluateRequestDTO {
  return {
    dishName: domain.dishName,
    description: domain.description,
  };
}

export function mapEvaluateResponseToDomain(dto: EvaluateResponseDTO): EvaluateResponseDomain {
  return {
    success: dto.success,
    score: dto.score,
    feedback: dto.feedback,
    creativityScore: dto.creativityScore,
    techniqueScore: dto.techniqueScore,
    appealScore: dto.appealScore,
    title: dto.title,
    achievement: dto.achievement,
    error: dto.error,
  };
}
