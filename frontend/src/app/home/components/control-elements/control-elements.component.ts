import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'app-control-elements',
  templateUrl: './control-elements.component.html',
  styleUrls: ['./control-elements.component.scss']
})
export class ControlElementsComponent implements OnInit {

  public video = '-';
  public videoUrl = '';

  public loop = false;
  public autoplay = false;

  constructor() { }

  ngOnInit(): void {
  }

  open(): void {

  }

  updatePlayerConfig(): void {

  }
}
