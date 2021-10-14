import { ComponentFixture, TestBed } from '@angular/core/testing';

import { BottomButtonComponent } from './bottom-button.component';

describe('BottomButtonComponent', () => {
  let component: BottomButtonComponent;
  let fixture: ComponentFixture<BottomButtonComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ BottomButtonComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(BottomButtonComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
