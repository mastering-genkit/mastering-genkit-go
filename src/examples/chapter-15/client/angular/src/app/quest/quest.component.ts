import { Component, OnInit, OnDestroy } from '@angular/core';
import { CommonModule } from '@angular/common';
import { Observable, Subscription } from 'rxjs';
import { GameStateService, GenerateRecipeService, CreateImageService, EvaluateDishService } from '../services';
import { GameState, GameStep } from '../../domain/models/game';

@Component({
  selector: 'app-quest',
  standalone: true,
  imports: [CommonModule],
  templateUrl: './quest.component.html',
  styleUrl: './quest.component.css'
})
export class QuestComponent implements OnInit, OnDestroy {
  state$: Observable<GameState>;
  availableIngredients: Array<{name: string, emoji: string}> = [];
  GameStep = GameStep; // Expose enum to template
  
  private subscriptions = new Subscription();

  constructor(
    private gameStateService: GameStateService,
    private generateRecipeService: GenerateRecipeService,
    private createImageService: CreateImageService,
    private evaluateDishService: EvaluateDishService
  ) {
    this.state$ = this.gameStateService.state$;
    this.availableIngredients = this.gameStateService.availableIngredients;
  }

  ngOnInit(): void {
    // Subscribe to recipe stream
    this.subscriptions.add(
      this.generateRecipeService.recipeStream$.subscribe(response => {
        if (response.type === 'content' && response.content) {
          const currentState = this.gameStateService.currentState;
          const newRecipe = (currentState.recipe || '') + response.content;
          this.gameStateService.setRecipe(newRecipe);
        } else if (response.type === 'done') {
          // Recipe generation complete, wait before moving to image step
          setTimeout(() => {
            this.gameStateService.startImageGeneration();
            this.createDishImage();
          }, 1500);
        } else if (response.type === 'error') {
          this.gameStateService.setError(response.error || 'Recipe generation failed');
        }
      })
    );
  }

  ngOnDestroy(): void {
    this.subscriptions.unsubscribe();
  }

  startGame(): void {
    this.gameStateService.startGame();
  }

  toggleIngredient(ingredient: string): void {
    const state = this.gameStateService.currentState;
    if (state.selectedIngredients.includes(ingredient)) {
      this.gameStateService.removeIngredient(ingredient);
    } else {
      this.gameStateService.addIngredient(ingredient);
    }
  }

  async startRecipeGeneration(): Promise<void> {
    const state = this.gameStateService.currentState;
    if (state.selectedIngredients.length === 4) {
      this.gameStateService.startRecipeGeneration();
      await this.generateRecipeService.generateRecipe({
        ingredients: state.selectedIngredients
      });
    }
  }

  private async createDishImage(): Promise<void> {
    const state = this.gameStateService.currentState;
    if (state.recipe) {
      const dishName = this.extractDishName(state.recipe);
      
      // Simulate progress while generating
      let currentProgress = 0;
      const progressInterval = setInterval(() => {
        currentProgress += 15;
        if (currentProgress <= 90) {
          this.gameStateService.setImageProgress(currentProgress);
        }
      }, 300);
      
      try {
        const result = await this.createImageService.createImage({
          dishName,
          description: `A beautiful dish made with: ${state.selectedIngredients.join(', ')}`
        });
        
        clearInterval(progressInterval);
        
        if (result && result.success && result.imageUrl) {
          this.gameStateService.setImage(result.imageUrl);
          // Wait before moving to evaluation step
          setTimeout(() => {
            this.gameStateService.startEvaluation();
            setTimeout(() => {
              this.evaluateDishStep();
            }, 1000);
          }, 1000);
        } else {
          this.gameStateService.setError('Image generation failed');
        }
      } catch (error) {
        clearInterval(progressInterval);
        this.gameStateService.setError('Image generation failed');
      }
    }
  }

  private async evaluateDishStep(): Promise<void> {
    const state = this.gameStateService.currentState;
    if (state.recipe) {
      const dishName = this.extractDishName(state.recipe);
      const result = await this.evaluateDishService.evaluateDish({
        dishName,
        description: state.recipe
      });
      
      if (result && result.success) {
        this.gameStateService.setEvaluation({
          score: result.score || 0,
          feedback: result.feedback || '',
          title: result.title,
          achievement: result.achievement
        });
      } else {
        this.gameStateService.setError('Evaluation failed');
      }
    }
  }

  private extractDishName(recipe: string): string {
    // Simple extraction - look for title patterns or use ingredients
    const lines = recipe.split('\n');
    const titleLine = lines.find(line => 
      line.includes('#') || 
      line.toLowerCase().includes('recipe') ||
      line.toLowerCase().includes('dish')
    );
    
    if (titleLine) {
      return titleLine.replace(/#+\s*/, '').trim();
    }
    
    // Fallback to ingredient combination
    const state = this.gameStateService.currentState;
    return state.selectedIngredients.join(' ') + ' dish';
  }

  resetGame(): void {
    this.gameStateService.resetGame();
  }

  getStepIndex(step: GameStep): number {
    const stepOrder = [GameStep.Ready, GameStep.SelectIngredients, GameStep.Recipe, GameStep.Image, GameStep.Evaluation, GameStep.Result];
    return stepOrder.indexOf(step);
  }

  Math = Math; // Expose Math to template
}
