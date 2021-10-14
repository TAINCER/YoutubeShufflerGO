import { TestBed } from '@angular/core/testing';

import { WailsService } from './wails.service';

describe('WailsService', () => {
  let service: WailsService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(WailsService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
