// DTO for createImage flow response
export interface ImageResponseDTO {
  success: boolean;
  imageUrl?: string;
  dishName?: string;
  error?: string;
}
