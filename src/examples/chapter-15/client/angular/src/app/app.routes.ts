import { Routes } from '@angular/router';

export const routes: Routes = [
  { path: '', redirectTo: '/quest', pathMatch: 'full' },
  { path: 'quest', loadComponent: () => import('./quest/quest.component').then(m => m.QuestComponent) }
];
