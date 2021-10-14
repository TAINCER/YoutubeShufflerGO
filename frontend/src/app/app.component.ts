import { Component } from '@angular/core';

// FIXME: Not working with Wails beta yet
declare var window: any;

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.scss']
})
export class AppComponent {

  close(): void {
    window.runtime.Quit();
  }

  minimize(): void {
    window.runtime.WindowMinimise();
  }
}

