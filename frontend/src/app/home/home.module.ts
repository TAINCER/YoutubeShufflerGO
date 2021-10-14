import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { MatGridListModule } from '@angular/material/grid-list';
import { MatProgressSpinnerModule } from '@angular/material/progress-spinner';
import { MatCheckboxModule } from '@angular/material/checkbox';
import { MatButtonModule } from '@angular/material/button';
import { ClipboardModule } from '@angular/cdk/clipboard';

import { HomeRoutingModule } from './home-routing.module';

import { BottomButtonComponent } from './components/bottom-button/bottom-button.component';
import { ControlElementsComponent } from './components/control-elements/control-elements.component';
import { VideoViewComponent } from './components/video-view/video-view.component';
import { HomeComponent } from './pages/home/home.component';


@NgModule({
  declarations: [
    HomeComponent,
    BottomButtonComponent,
    ControlElementsComponent,
    VideoViewComponent
  ],
  imports: [
    CommonModule,
    HomeRoutingModule,
    MatGridListModule,
    MatButtonModule,
    MatProgressSpinnerModule,
    MatCheckboxModule,
    ClipboardModule
  ]
})
export class HomeModule { }
