import { Injectable } from '@angular/core';

declare var window: any;

@Injectable({
  providedIn: 'root'
})
export class WailsService {

  constructor() { }

  async call(method: Promise<any>): Promise<any> {
    return await new Promise((resolve, reject) => {
      method.then((res: any) => resolve(res)).catch((err: any) => reject(err))
    });
  }
}
