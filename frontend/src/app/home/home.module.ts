import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';

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
    HomeRoutingModule
  ]
})
export class HomeModule { }
