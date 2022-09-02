import {Injectable, NestMiddleware} from '@nestjs/common';
import path from 'path';

@Injectable()
export class FrontendMiddleware implements NestMiddleware {
    use(req: any, res: any, next: (error?: any) => void) {
        throw new Error('Method not implemented.');
    }
}

